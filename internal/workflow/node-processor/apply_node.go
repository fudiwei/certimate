package nodeprocessor

import (
	"context"

	"github.com/usual2970/certimate/internal/applicant"
	"github.com/usual2970/certimate/internal/domain"
	"github.com/usual2970/certimate/internal/pkg/utils/x509"
	"github.com/usual2970/certimate/internal/repository"
	"github.com/usual2970/certimate/internal/utils/xtime"
)

type applyNode struct {
	node       *domain.WorkflowNode
	outputRepo WorkflowOutputRepository
	*Logger
}

func NewApplyNode(node *domain.WorkflowNode) *applyNode {
	return &applyNode{
		node:       node,
		Logger:     NewLogger(node),
		outputRepo: repository.NewWorkflowOutputRepository(),
	}
}

type WorkflowOutputRepository interface {
	// 查询节点输出
	Get(ctx context.Context, nodeId string) (*domain.WorkflowOutput, error)

	// 保存节点输出
	Save(ctx context.Context, output *domain.WorkflowOutput, certificate *domain.Certificate, cb func(id string) error) error
}

// 申请节点根据申请类型执行不同的操作
func (a *applyNode) Run(ctx context.Context) error {
	a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "开始执行")
	// 查询是否申请过，已申请过则直接返回（先保持和 v0.2 一致）
	output, err := a.outputRepo.Get(ctx, a.node.Id)
	if err != nil && !domain.IsRecordNotFound(err) {
		a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "查询申请记录失败", err.Error())
		return err
	}

	if output != nil && output.Succeed {
		a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "已申请过")
		return nil
	}

	// 获取Applicant
	apply, err := applicant.GetWithApplyNode(a.node)
	if err != nil {
		a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "获取申请对象失败", err.Error())
		return err
	}

	// 申请
	certificate, err := apply.Apply()
	if err != nil {
		a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "申请失败", err.Error())
		return err
	}
	a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "申请成功")

	// 记录申请结果
	output = &domain.WorkflowOutput{
		Workflow: GetWorkflowId(ctx),
		NodeId:   a.node.Id,
		Node:     a.node,
		Succeed:  true,
	}

	cert, err := x509.ParseCertificateFromPEM(certificate.Certificate)
	if err != nil {
		a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "解析证书失败", err.Error())
		return err
	}

	certificateRecord := &domain.Certificate{
		SAN:               cert.Subject.CommonName,
		Certificate:       certificate.Certificate,
		PrivateKey:        certificate.PrivateKey,
		IssuerCertificate: certificate.IssuerCertificate,
		CertUrl:           certificate.CertUrl,
		CertStableUrl:     certificate.CertStableUrl,
		ExpireAt:          cert.NotAfter,
	}

	if err := a.outputRepo.Save(ctx, output, certificateRecord, func(id string) error {
		if certificateRecord != nil {
			certificateRecord.Id = id
		}

		return nil
	}); err != nil {
		a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "保存申请记录失败", err.Error())
		return err
	}

	a.AddOutput(ctx, xtime.BeijingTimeStr(), a.node.Name, "保存申请记录成功")

	return nil
}