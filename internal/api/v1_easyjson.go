// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	game "github.com/scribble-rs/scribble.rs/internal/game"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi(in *jlexer.Lexer, out *LobbyEntry) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "lobbyId":
			out.LobbyID = string(in.String())
		case "playerCount":
			out.PlayerCount = int(in.Int())
		case "maxPlayers":
			out.MaxPlayers = int(in.Int())
		case "round":
			out.Round = int(in.Int())
		case "rounds":
			out.Rounds = int(in.Int())
		case "drawingTime":
			out.DrawingTime = int(in.Int())
		case "customWords":
			out.CustomWords = bool(in.Bool())
		case "votekick":
			out.Votekick = bool(in.Bool())
		case "maxClientsPerIp":
			out.MaxClientsPerIP = int(in.Int())
		case "wordpack":
			out.Wordpack = string(in.String())
		case "state":
			out.State = game.State(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi(out *jwriter.Writer, in LobbyEntry) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"lobbyId\":"
		out.RawString(prefix[1:])
		out.String(string(in.LobbyID))
	}
	{
		const prefix string = ",\"playerCount\":"
		out.RawString(prefix)
		out.Int(int(in.PlayerCount))
	}
	{
		const prefix string = ",\"maxPlayers\":"
		out.RawString(prefix)
		out.Int(int(in.MaxPlayers))
	}
	{
		const prefix string = ",\"round\":"
		out.RawString(prefix)
		out.Int(int(in.Round))
	}
	{
		const prefix string = ",\"rounds\":"
		out.RawString(prefix)
		out.Int(int(in.Rounds))
	}
	{
		const prefix string = ",\"drawingTime\":"
		out.RawString(prefix)
		out.Int(int(in.DrawingTime))
	}
	{
		const prefix string = ",\"customWords\":"
		out.RawString(prefix)
		out.Bool(bool(in.CustomWords))
	}
	{
		const prefix string = ",\"votekick\":"
		out.RawString(prefix)
		out.Bool(bool(in.Votekick))
	}
	{
		const prefix string = ",\"maxClientsPerIp\":"
		out.RawString(prefix)
		out.Int(int(in.MaxClientsPerIP))
	}
	{
		const prefix string = ",\"wordpack\":"
		out.RawString(prefix)
		out.String(string(in.Wordpack))
	}
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix)
		out.String(string(in.State))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LobbyEntry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LobbyEntry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LobbyEntry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LobbyEntry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi(l, v)
}
func easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi1(in *jlexer.Lexer, out *LobbyEntries) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(LobbyEntries, 0, 8)
			} else {
				*out = LobbyEntries{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *LobbyEntry
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(LobbyEntry)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi1(out *jwriter.Writer, in LobbyEntries) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v LobbyEntries) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LobbyEntries) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LobbyEntries) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LobbyEntries) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi1(l, v)
}
func easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi2(in *jlexer.Lexer, out *LobbyData) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "lobbyId":
			out.LobbyID = string(in.String())
		case "drawingBoardBaseWidth":
			out.DrawingBoardBaseWidth = int(in.Int())
		case "drawingBoardBaseHeight":
			out.DrawingBoardBaseHeight = int(in.Int())
		case "minBrushSize":
			out.MinBrushSize = int(in.Int())
		case "maxBrushSize":
			out.MaxBrushSize = int(in.Int())
		case "canvasColor":
			(out.CanvasColor).UnmarshalEasyJSON(in)
		case "suggestedBrushSizes":
			if in.IsNull() {
				in.Skip()
			} else {
				copy(out.SuggestedBrushSizes[:], in.Bytes())
			}
		case "maxPlayers":
			out.MaxPlayers = int(in.Int())
		case "public":
			out.Public = bool(in.Bool())
		case "enableVotekick":
			out.EnableVotekick = bool(in.Bool())
		case "customWordsChance":
			out.CustomWordsChance = int(in.Int())
		case "clientsPerIpLimit":
			out.ClientsPerIPLimit = int(in.Int())
		case "drawingTime":
			out.DrawingTime = int(in.Int())
		case "rounds":
			out.Rounds = int(in.Int())
		case "minDrawingTime":
			out.MinDrawingTime = int64(in.Int64())
		case "maxDrawingTime":
			out.MaxDrawingTime = int64(in.Int64())
		case "minRounds":
			out.MinRounds = int64(in.Int64())
		case "maxRounds":
			out.MaxRounds = int64(in.Int64())
		case "minMaxPlayers":
			out.MinMaxPlayers = int64(in.Int64())
		case "maxMaxPlayers":
			out.MaxMaxPlayers = int64(in.Int64())
		case "minClientsPerIpLimit":
			out.MinClientsPerIPLimit = int64(in.Int64())
		case "maxClientsPerIpLimit":
			out.MaxClientsPerIPLimit = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi2(out *jwriter.Writer, in LobbyData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"lobbyId\":"
		out.RawString(prefix[1:])
		out.String(string(in.LobbyID))
	}
	{
		const prefix string = ",\"drawingBoardBaseWidth\":"
		out.RawString(prefix)
		out.Int(int(in.DrawingBoardBaseWidth))
	}
	{
		const prefix string = ",\"drawingBoardBaseHeight\":"
		out.RawString(prefix)
		out.Int(int(in.DrawingBoardBaseHeight))
	}
	{
		const prefix string = ",\"minBrushSize\":"
		out.RawString(prefix)
		out.Int(int(in.MinBrushSize))
	}
	{
		const prefix string = ",\"maxBrushSize\":"
		out.RawString(prefix)
		out.Int(int(in.MaxBrushSize))
	}
	{
		const prefix string = ",\"canvasColor\":"
		out.RawString(prefix)
		(in.CanvasColor).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"suggestedBrushSizes\":"
		out.RawString(prefix)
		out.Base64Bytes(in.SuggestedBrushSizes[:])
	}
	{
		const prefix string = ",\"maxPlayers\":"
		out.RawString(prefix)
		out.Int(int(in.MaxPlayers))
	}
	{
		const prefix string = ",\"public\":"
		out.RawString(prefix)
		out.Bool(bool(in.Public))
	}
	{
		const prefix string = ",\"enableVotekick\":"
		out.RawString(prefix)
		out.Bool(bool(in.EnableVotekick))
	}
	{
		const prefix string = ",\"customWordsChance\":"
		out.RawString(prefix)
		out.Int(int(in.CustomWordsChance))
	}
	{
		const prefix string = ",\"clientsPerIpLimit\":"
		out.RawString(prefix)
		out.Int(int(in.ClientsPerIPLimit))
	}
	{
		const prefix string = ",\"drawingTime\":"
		out.RawString(prefix)
		out.Int(int(in.DrawingTime))
	}
	{
		const prefix string = ",\"rounds\":"
		out.RawString(prefix)
		out.Int(int(in.Rounds))
	}
	{
		const prefix string = ",\"minDrawingTime\":"
		out.RawString(prefix)
		out.Int64(int64(in.MinDrawingTime))
	}
	{
		const prefix string = ",\"maxDrawingTime\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxDrawingTime))
	}
	{
		const prefix string = ",\"minRounds\":"
		out.RawString(prefix)
		out.Int64(int64(in.MinRounds))
	}
	{
		const prefix string = ",\"maxRounds\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxRounds))
	}
	{
		const prefix string = ",\"minMaxPlayers\":"
		out.RawString(prefix)
		out.Int64(int64(in.MinMaxPlayers))
	}
	{
		const prefix string = ",\"maxMaxPlayers\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxMaxPlayers))
	}
	{
		const prefix string = ",\"minClientsPerIpLimit\":"
		out.RawString(prefix)
		out.Int64(int64(in.MinClientsPerIPLimit))
	}
	{
		const prefix string = ",\"maxClientsPerIpLimit\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxClientsPerIPLimit))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LobbyData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LobbyData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson102f8a2fEncodeGithubComScribbleRsScribbleRsInternalApi2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LobbyData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LobbyData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson102f8a2fDecodeGithubComScribbleRsScribbleRsInternalApi2(l, v)
}
