package tmpl

import (
    "text/template"
    "bytes"

    "github.com/k0tletka/minecraft-reserve-copy/common/types"
)

func GenerateArchiveName(archiveNameTemplate string, templateData *types.ArchiveTemplateData) (string, error) {
    // Create template and generate archive name
    archiveNameTemplateObj, err := template.New("archive-name-template").Parse(archiveNameTemplate)
    if err != nil {
        return "", err
    }

    archiveNameBuffer := &bytes.Buffer{}
    if err := archiveNameTemplateObj.Execute(archiveNameBuffer, templateData); err != nil {
        return "", err
    }

    return archiveNameBuffer.String(), nil
}
