{{- $table_alias := .Aliases.Table .Table.Name -}}
{{- $table_name := .Table.Name -}}

// {{$table_alias.UpSingular}} is {{$table_alias.UpSingular}} entity object.
type {{$table_alias.UpSingular}} struct {
	{{- range $column := .Table.Columns -}}
	{{- $col_alias := $table_alias.Column $column.Name -}}
	{{- $col_name := $column.Name -}}
	{{- range $column.Comment | splitLines -}} // {{ . }}
	{{end -}}
	{{if ignore $table_name $col_name $.TagIgnore -}}
	{{$col_alias}} {{$column.Type}} `{{generateIgnoreTags $.Tags}}db:"{{$column.Name}}" json:"-"`
	{{else if eq $.StructTagCasing "title" -}}
	{{$col_alias}} {{$column.Type}} `{{generateTags $.Tags $column.Name}}json:"{{$column.Name | titleCase}}{{if $column.Nullable}},omitempty{{end}}"`
	{{else if eq $.StructTagCasing "camel" -}}
	{{$col_alias}} {{$column.Type}} `{{generateTags $.Tags $column.Name}}json:"{{$column.Name | camelCase}}{{if $column.Nullable}},omitempty{{end}}"`
	{{else if eq $.StructTagCasing "alias" -}}
	{{$col_alias}} {{$column.Type}} `{{generateTags $.Tags $col_alias}}json:"{{$col_alias}}{{if $column.Nullable}},omitempty{{end}}"`
	{{else -}}
	{{$col_alias}} {{$column.Type}} `{{generateTags $.Tags $column.Name}}json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
	{{end -}}
	{{end -}}
	{{- if .Table.IsJoinTable -}}
	{{- else}}
	{{end -}}
}

{{range $column := .Table.Columns -}}
{{- $col_alias := $table_alias.Column $column.Name -}}
// Get{{$col_alias}} is get {{$column.Name}} value, if receiver is nil, returns the specified value.
func (m *{{$table_alias.UpSingular}}) Get{{$col_alias}}() {{$column.Type}} {
    if m == nil {
        {{if or (eq $column.Type "uint64") (eq $column.Type "int64") (eq $column.Type "uint") (eq $column.Type "int") (eq $column.Type "uint32") (eq $column.Type "int32") (eq $column.Type "uint16") (eq $column.Type "int16") (eq $column.Type "uint8") (eq $column.Type "int8") -}}
        return 0
        {{else if or (eq $column.Type "bool") -}}
        return false
        {{else if or (eq $column.Type "string") -}}
        return ""
        {{else if or (eq $column.Type "[]byte") -}}
        return nil
        {{else if or (eq $column.Type "decimal.Decimal") -}}
        return decimal.Zero()
        {{else if or (eq $column.Type "times.Date") (eq $column.Type "times.Timestamp") -}}
        return {{$column.Type}}{}
    	{{else -}}
        return {{$column.Type}}(0)
        {{end -}}
    }
    return m.{{$col_alias}}
}
{{end -}}
