import type { Metadata } from "next";
import localFont from "next/font/local";
import Header from "@/features/Header";

import { ClerkProvider } from "@clerk/nextjs";

import "./globals.css";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

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
        <body
          className={`${geistSans.variable} ${geistMono.variable} antialiased`}
        >
          <Header />
          <div className="flex justify-center w-full">
            <main className="flex max-w-7xl w-full">{children}</main>
          </div>
        </body>
      </html>
    </ClerkProvider>
  );
}
