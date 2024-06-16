// Code generated by ent, DO NOT EDIT.

package waterlog

import (
	"plant-watering/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldLTE(FieldID, id))
}

// Seconds applies equality check predicate on the "seconds" field. It's identical to SecondsEQ.
func Seconds(v int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldSeconds, v))
}

// Channel applies equality check predicate on the "channel" field. It's identical to ChannelEQ.
func Channel(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldChannel, v))
}

// Manual applies equality check predicate on the "manual" field. It's identical to ManualEQ.
func Manual(v bool) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldManual, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldTime, v))
}

// SecondsEQ applies the EQ predicate on the "seconds" field.
func SecondsEQ(v int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldSeconds, v))
}

// SecondsNEQ applies the NEQ predicate on the "seconds" field.
func SecondsNEQ(v int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNEQ(FieldSeconds, v))
}

// SecondsIn applies the In predicate on the "seconds" field.
func SecondsIn(vs ...int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldIn(FieldSeconds, vs...))
}

// SecondsNotIn applies the NotIn predicate on the "seconds" field.
func SecondsNotIn(vs ...int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNotIn(FieldSeconds, vs...))
}

// SecondsGT applies the GT predicate on the "seconds" field.
func SecondsGT(v int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldGT(FieldSeconds, v))
}

// SecondsGTE applies the GTE predicate on the "seconds" field.
func SecondsGTE(v int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldGTE(FieldSeconds, v))
}

// SecondsLT applies the LT predicate on the "seconds" field.
func SecondsLT(v int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldLT(FieldSeconds, v))
}

// SecondsLTE applies the LTE predicate on the "seconds" field.
func SecondsLTE(v int) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldLTE(FieldSeconds, v))
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldChannel, v))
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNEQ(FieldChannel, v))
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldIn(FieldChannel, vs...))
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNotIn(FieldChannel, vs...))
}

// ChannelGT applies the GT predicate on the "channel" field.
func ChannelGT(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldGT(FieldChannel, v))
}

// ChannelGTE applies the GTE predicate on the "channel" field.
func ChannelGTE(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldGTE(FieldChannel, v))
}

// ChannelLT applies the LT predicate on the "channel" field.
func ChannelLT(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldLT(FieldChannel, v))
}

// ChannelLTE applies the LTE predicate on the "channel" field.
func ChannelLTE(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldLTE(FieldChannel, v))
}

// ChannelContains applies the Contains predicate on the "channel" field.
func ChannelContains(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldContains(FieldChannel, v))
}

// ChannelHasPrefix applies the HasPrefix predicate on the "channel" field.
func ChannelHasPrefix(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldHasPrefix(FieldChannel, v))
}

// ChannelHasSuffix applies the HasSuffix predicate on the "channel" field.
func ChannelHasSuffix(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldHasSuffix(FieldChannel, v))
}

// ChannelEqualFold applies the EqualFold predicate on the "channel" field.
func ChannelEqualFold(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEqualFold(FieldChannel, v))
}

// ChannelContainsFold applies the ContainsFold predicate on the "channel" field.
func ChannelContainsFold(v string) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldContainsFold(FieldChannel, v))
}

// ManualEQ applies the EQ predicate on the "manual" field.
func ManualEQ(v bool) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldManual, v))
}

// ManualNEQ applies the NEQ predicate on the "manual" field.
func ManualNEQ(v bool) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNEQ(FieldManual, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v time.Time) predicate.WaterLog {
	return predicate.WaterLog(sql.FieldLTE(FieldTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WaterLog) predicate.WaterLog {
	return predicate.WaterLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WaterLog) predicate.WaterLog {
	return predicate.WaterLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WaterLog) predicate.WaterLog {
	return predicate.WaterLog(sql.NotPredicates(p))
}