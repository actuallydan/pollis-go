import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Header from "@/features/Header";
import { ClerkProvider } from "@clerk/nextjs";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Pollis",
  description: "Real connections to real communities",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <ClerkProvider>
      <html lang="en">
        <body className={inter.className}>
          <Header />
          <div className="flex justify-center w-full">
            <main className="flex max-w-7xl w-full">{children}</main>
          </div>
        </body>
      </html>
    </ClerkProvider>
  );
}
