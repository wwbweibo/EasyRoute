package Route

type Pipeline struct {
	handlerList []*PipelineHandler
}

type PipelineHandler interface {
	Handle(c RequestContext, pipeline *Pipeline)
}

func (receiver *Pipeline) Next(c *RequestContext) {

}
