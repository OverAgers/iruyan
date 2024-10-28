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
      <body>
        {children}
      </body>
    </html>
  );
}
