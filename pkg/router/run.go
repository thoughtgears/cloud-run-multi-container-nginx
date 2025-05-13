package router

import "fmt"

func (r *Router) Run() error {
	if err := r.Engine.Run(":" + r.port); err != nil {
		return fmt.Errorf("error running router: %w", err)
	}

	return nil
}
