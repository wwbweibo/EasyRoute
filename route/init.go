package route

import (
	types2 "github.com/wwbweibo/EasyRoute/internal/types"
)

var typeCollectionInstance = types2.NewTypeCollect()

func InjectTypes(collect *types2.TypeCollect) {
	typeCollectionInstance = collect
}
