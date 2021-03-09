package utils 

import "regexp"

var UserMention = regexp.MustCompile(`^<@!?(\d{17,19})>$`)

var UserID = regexp.MustCompile(`^(\d{17,19})$`)

var ChannelMention = regexp.MustCompile(`^<#(\d{17,19})>$`)

var ChannelID = regexp.MustCompile(`^(\d{17,19})$`)

var Symbols = regexp.MustCompile("[#@!<>]")