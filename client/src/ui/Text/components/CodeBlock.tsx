import { CodeComponent } from "react-markdown/lib/ast-to-react";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { a11yDark } from "react-syntax-highlighter/dist/cjs/styles/prism";

const customCode: CodeComponent = ({ inline, className, children }) => {
  const style = a11yDark;
  const match = /language-(\w+)(:?.+)?/.exec(className || "");
  const lang = match?.[1] ? match[1] : "";
  const name = match?.[2] ? match[2].slice(1) : "";
  return !inline && match ? (
    <>
      {name && <span className="px-2 py-1 text-xs">{name}</span>}
      <SyntaxHighlighter
        style={style}
        language={lang}
        PreTag="div"
        // eslint-disable-next-line tailwindcss/no-custom-classname
        className="md-codeblock">
        {String(children).replace(/\n$/, "")}
      </SyntaxHighlighter>
    </>
  ) : (
    <code className="text-xs text-red-400">{children}</code>
  );
};
// codeタグで記述された部分はcustomCodeに置き換えられる
export const CodeBlock = {
  code: customCode,
};
