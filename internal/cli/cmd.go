package cli

import (
	"github.com/AEKDA/aebuilder/internal/builder"
	"github.com/spf13/cobra"
)

type Cmd struct {
	root *cobra.Command
}

func New() *Cmd {
	var (
		source, output, name string
		tags                 []string
	)

	root := &cobra.Command{
		Use:   "aebuilder",
		Short: "aebuilder TODO!",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	genCmd := &cobra.Command{
		Use:   "gen",
		Short: "Генерирует модель и билдер для нее из прото",
		Long:  `Генерирует модель и билдер для нее из прото`,
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := builder.New(source, output, name, tags)
			if err != nil {
				return err
			}
			return app.Run()
		},
	}

	genCmd.Flags().StringVarP(&source, "source", "s", "", "укажите путь до прото файла (обязательно)")
	genCmd.Flags().StringVarP(&name, "name", "n", "", "укажите имя структуры для которой надо сгенерировать билдер и модель (обязательно)")
	genCmd.Flags().StringVarP(&output, "output", "o", "", "укажите путь до места куда надо сохранить сгенерированные файлы (обязательно)")
	genCmd.Flags().StringSliceVar(&tags, "tag", []string{}, "Укажите теги которые будут в итоговой модели в следующем формате name:case")

	genCmd.MarkFlagRequired("source")
	genCmd.MarkFlagRequired("output")
	genCmd.MarkFlagRequired("name")

	root.AddCommand(genCmd)

	return &Cmd{
		root: root,
	}
}

func (r *Cmd) Run() error {
	return r.root.Execute()
}
