package app

import content "promotion/internal/content"

type Modules struct {
	Content *content.Module
}

func initModules(infra *infrastructure) *Modules {
	return &Modules{
		Content: content.NewModule(infra.log, infra.db.MySQL),
	}
}
