package arachne

import (
	//	"archive/zip"
	"context"
)

type ProtectedEntityTypeManager interface {
	GetTypeName() string
	GetProtectedEntity(ctx context.Context, id ProtectedEntityID) (ProtectedEntity, error)
	GetProtectedEntities(ctx context.Context) ([]ProtectedEntity, error)
	Copy(ctx context.Context, pe ProtectedEntity) (ProtectedEntity, error)
	CopyFromInfo(ctx context.Context, info ProtectedEntityInfo) (ProtectedEntity, error)
	//Serialize(pe ProtectedEntity, out Zip.Writer)
	//Deserialize(is ZipInputStream, ProtectedEntityInfo peInfo) ProtectedEntity
	//SerializeData(pe ProtectedEntity, out OutputStream)
}
