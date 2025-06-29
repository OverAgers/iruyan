import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "iruyan",
  description: "オンライン自習室〜居る家ん〜",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body style={{ boxSizing: "border-box", margin: 0, padding: 0 }}>
        {children}
      </body>
    </html>
  );
}
