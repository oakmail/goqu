package sqlite3

import "github.com/trtlio/goqu"

var (
	placeholder_rune    = '?'
	quote_rune          = '`'
	singlq_quote        = '\''
	default_values_frag = []byte("")
	sqlite3_true        = []byte("1")
	sqlite3_false       = []byte("0")
	time_format         = "2006-01-02 15:04:05"
	operator_lookup     = map[goqu.BooleanOperation][]byte{
		goqu.EQ_OP:                []byte("="),
		goqu.NEQ_OP:               []byte("!="),
		goqu.GT_OP:                []byte(">"),
		goqu.GTE_OP:               []byte(">="),
		goqu.LT_OP:                []byte("<"),
		goqu.LTE_OP:               []byte("<="),
		goqu.IN_OP:                []byte("IN"),
		goqu.NOT_IN_OP:            []byte("NOT IN"),
		goqu.IS_OP:                []byte("IS"),
		goqu.IS_NOT_OP:            []byte("IS NOT"),
		goqu.LIKE_OP:              []byte("LIKE"),
		goqu.NOT_LIKE_OP:          []byte("NOT LIKE"),
		goqu.I_LIKE_OP:            []byte("LIKE"),
		goqu.NOT_I_LIKE_OP:        []byte("NOT LIKE"),
		goqu.REGEXP_LIKE_OP:       []byte("REGEXP"),
		goqu.REGEXP_NOT_LIKE_OP:   []byte("NOT REGEXP"),
		goqu.REGEXP_I_LIKE_OP:     []byte("REGEXP"),
		goqu.REGEXP_NOT_I_LIKE_OP: []byte("NOT REGEXP"),
	}
	escape_runes = map[rune][]byte{
		'\'': []byte("\\'"),
		'"':  []byte("\\\""),
		'\\': []byte("\\\\"),
		'\n': []byte("\\n"),
		'\r': []byte("\\r"),
		0:    []byte("\\x00"),
		0x1a: []byte("\\x1a"),
	}
	row_id_column_name = "rowid"
)

type DatasetAdapter struct {
	*goqu.DefaultAdapter
}

func (me *DatasetAdapter) SupportsReturn() bool {
	return false
}

func (me *DatasetAdapter) SupportsLimitOnDelete() bool {
	return true
}

func (me *DatasetAdapter) SupportsLimitOnUpdate() bool {
	return true
}

func (me *DatasetAdapter) SupportsOrderByOnDelete() bool {
	return true
}

func (me *DatasetAdapter) SupportsOrderByOnUpdate() bool {
	return true
}

func newDatasetAdapter(ds *goqu.Dataset) goqu.Adapter {
	def := goqu.NewDefaultAdapter(ds).(*goqu.DefaultAdapter)
	def.PlaceHolderRune = placeholder_rune
	def.IncludePlaceholderNum = false
	def.QuoteRune = quote_rune
	def.DefaultValuesFragment = default_values_frag
	def.True = sqlite3_true
	def.False = sqlite3_false
	def.TimeFormat = time_format
	def.BooleanOperatorLookup = operator_lookup
	def.UseLiteralIsBools = false
	def.EscapedRunes = escape_runes
	def.RowIDColumnName = row_id_column_name
	return &DatasetAdapter{def}
}

func init() {
	goqu.RegisterAdapter("sqlite3", newDatasetAdapter)
}
