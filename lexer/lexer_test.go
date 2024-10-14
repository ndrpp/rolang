package lexer

import (
	"testing"

	"rolang/token"
)

func TestNextToken(t *testing.T) {
	input := `var five = 5
    var ten = 10
    var add = fn(x, y) {
        x + y
    }
    var result = add(five, ten)
    !-/*5
    5 < 10 > 5
    dacă 5 < 10 {
        returnează adevărat
    } altfel {
        returnează fals
    }
    10 == 10
    10 != 9
    "foobar"
    "foo bar"
    [1, 2]
    {"foo": "bar"}
    pentru user în users {
        afișează(user.name)
    }
    i++
    j--
    k += 1
    l -= 5
    5 - 1
    !adevărat
    `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.VAR, "var"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.VAR, "var"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.IF, "dacă"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.RETURN, "returnează"},
		{token.TRUE, "adevărat"},
		{token.RBRACE, "}"},
		{token.ELSE, "altfel"},
		{token.LBRACE, "{"},
		{token.RETURN, "returnează"},
		{token.FALSE, "fals"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.FOR, "pentru"},
		{token.IDENT, "user"},
		{token.IN, "în"},
		{token.IDENT, "users"},
		{token.LBRACE, "{"},
		{token.PRINT, "afișează"},
		{token.LPAREN, "("},
		{token.IDENT, "user"},
		{token.DOT, "."},
		{token.IDENT, "name"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.IDENT, "i"},
		{token.INCREMENT, "++"},
		{token.IDENT, "j"},
		{token.DECREMENT, "--"},
		{token.IDENT, "k"},
		{token.PLUS_EQ, "+="},
		{token.IDENT, "l"},
		{token.MINUS_EQ, "-="},
		{token.INT, "5"},
		{token.MINUS, "-"},
		{token.BANG, "!"},
		{token.TRUE, "adevărat"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
