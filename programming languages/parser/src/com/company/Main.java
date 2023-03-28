package com.company;

/*

<Decl>::= <Type> <Ptr>
<Type>::= int | float
<Ptr>::= *<Ptr> | <Prim> <Dims>
<Dims>::= [ NUMBER ] <Dims> | ε
<Prim>::= IDENT | ( <Ptr> )

 */

import java.util.ArrayList;
import java.util.List;

public class Main {

    public static void main(String[] args) {
        String str = "int *(*a[10][20])[5]";
        // <Type> * ( * <Prim> <Dims> <Dims> ) <Dims>
        // <Type> * ( * <Ptr> ) <Dims>
        // <Type> * ( <Ptr> ) <Dims>
        // <Type> * <Prim> <Dims>
        // <Type> * <Ptr>
        // <Type> <Ptr>
        // <Decl>
        System.out.println("str = " + str);
        System.out.println(Decl(new LexemeBuffer(LexAnalyze(str))));
    }

    public enum LexemeType {
        LEFT_BRACKET, RIGHT_BRACKET,
        LEFT_SQUARE_BRACKET, RIGHT_SQUARE_BRACKET,
        TYPE, NUMBER, IDENT, STAR,
        EOF
    }

    public static class Lexeme {
        LexemeType type;
        String value;
        int line, column;

        public Lexeme (LexemeType type, String value, int line, int column) {
            this.type = type;
            this.value = value;
            this.line = line;
            this.column = column;
        }

        public Lexeme (LexemeType type, Character value, int line, int column) {
            this.type = type;
            this.value = value.toString();
            this.line = line;
            this.column = column;
        }

        @Override
        public String toString() {
            return "Lexeme{" +
                    "type=" + type +
                    ", value='" + value + '\'' +
                    ", line=" + line +
                    ", column=" + column +
                    '}';
        }
    }

    public static class LexemeBuffer {
        private int pos;
        public List<Lexeme> lexemes;

        LexemeBuffer(List<Lexeme> lexemes) {
            this.lexemes = lexemes;
        }

        public Lexeme next() {
            return lexemes.get(pos++);
        }

        public void back() {
            pos--;
        }

        public Lexeme current() {
            return lexemes.get(pos);
        }
    }

    public static List<Lexeme> LexAnalyze(String expText) {
        ArrayList<Lexeme> lexemes = new ArrayList<>();
        int pos = 0;
        int line = 0, column = 0;
        while (pos < expText.length()) {
            char c = expText.charAt(pos);
            switch (c) {
                case '(':
                    lexemes.add(new Lexeme(LexemeType.LEFT_BRACKET, c, line, column));
                    pos++;
                    column++;
                    continue;
                case ')':
                    lexemes.add(new Lexeme(LexemeType.RIGHT_BRACKET, c, line, column));
                    pos++;
                    column++;
                    continue;
                case '[':
                    lexemes.add(new Lexeme(LexemeType.LEFT_SQUARE_BRACKET, c, line, column));
                    pos++;
                    column++;
                    continue;
                case ']':
                    lexemes.add(new Lexeme(LexemeType.RIGHT_SQUARE_BRACKET, c, line, column));
                    pos++;
                    column++;
                    continue;
                case '*':
                    lexemes.add(new Lexeme(LexemeType.STAR, c, line, column));
                    pos++;
                    column++;
                    continue;
                default:
                    if(c <= '9' && c >= '0') {
                        StringBuilder sb = new StringBuilder();
                        do {
                            sb.append(c);
                            pos++;
                            column++;
                            if(pos >= expText.length()) break;
                            c = expText.charAt(pos);
                        } while(c <= '9' && c >= '0');
                        lexemes.add(new Lexeme(LexemeType.NUMBER, sb.toString(), line, column));
                        continue;
                    }
                    if(c <= 'Z' && c >= 'A' || c >= 'a' && c <= 'z') {
                        StringBuilder sb = new StringBuilder();
                        do {
                            sb.append(c);
                            pos++;
                            column++;
                            if(pos >= expText.length()) break;
                            c = expText.charAt(pos);
                        } while (c <= 'Z' && c >= 'A' || c >= 'a' && c <= 'z' || c <= '9' && c >= '0');
                        if(sb.toString().equals("int") || sb.toString().equals("float")) {
                            lexemes.add(new Lexeme(LexemeType.TYPE, sb.toString(), line, column));
                            continue;
                        } else {
                            lexemes.add(new Lexeme(LexemeType.IDENT, sb.toString(), line, column));
                            continue;
                        }
                    }
                    if(c == ' ') {
                        pos++;
                        column++;
                        continue;
                    }
                    if(c == '\n') {
                        pos++;
                        column = 0;
                        line++;
                    }
                    throw new RuntimeException("Unexpected character: " + c + "\nline: "
                            + (line + 1) + " column: " + (column + 1));
            }
        }
        lexemes.add(new Lexeme(LexemeType.EOF, "", line, column));
        return lexemes;
    }

/*

<Decl>::= <Type> <Ptr>
<Type>::= int | float
<Ptr>::= *<Ptr> | <Prim> <Dims>
<Dims>::= [ NUMBER ] <Dims> | ε
<Prim>::= IDENT | ( <Ptr> )

 */

    public static String Decl(LexemeBuffer lexemes) {
        return Type(lexemes) + Ptr(lexemes);
    }

    public static String Type(LexemeBuffer lexemes) {
        Lexeme lexeme = lexemes.next();
        if(lexeme.type == LexemeType.TYPE) {
            return "<Type> ";
        }
        throw new RuntimeException("syntax error at (" + lexeme.line
                + ", " + lexeme.column + ")");
    }

    public static String Ptr(LexemeBuffer lexemes) {
        Lexeme lexeme = lexemes.next();
        if (lexeme.type == LexemeType.STAR) {
            return "* " + Ptr(lexemes);
        }
        if(lexeme.type == LexemeType.IDENT || lexeme.type == LexemeType.LEFT_BRACKET) {
            lexemes.back();
            return Prim(lexemes) + Dims(lexemes);
        }
        throw new RuntimeException("syntax error at (" + lexeme.line
                + ", " + lexeme.column + ")");
    }

    public static String Dims(LexemeBuffer lexemes) {
        Lexeme lexeme = lexemes.next();
        if(lexeme.type == LexemeType.LEFT_SQUARE_BRACKET) {
            lexeme = lexemes.next();
            if(lexeme.type == LexemeType.NUMBER) {
                lexeme = lexemes.next();
                if(lexeme.type == LexemeType.RIGHT_SQUARE_BRACKET) {
                    return "<Dims> " + Dims(lexemes);
                }
            }
        }
        lexemes.back();
        return "";
    }

    public static String Prim(LexemeBuffer lexemes) {
        Lexeme lexeme = lexemes.next();
        switch (lexeme.type) {
            case IDENT:
                return "<Prim> ";
            case LEFT_BRACKET:
                String ptr = Ptr(lexemes);
                if(lexemes.next().type != LexemeType.RIGHT_BRACKET) {
                    lexemes.back();
                    throw new RuntimeException("syntax error at (" + lexemes.current().line
                            + ", " + lexemes.current().column + ")");
                }
                return "( " + ptr + ") ";
            default:
                throw new RuntimeException("syntax error at (" + lexeme.line
                        + ", " + lexeme.column + ")");
        }
    }

}
