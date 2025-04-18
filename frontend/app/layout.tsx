import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
    title: 'ConnectRPC Calculator',
    description: 'A calculator using ConnectRPC',
}

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
        <body className={inter.className}>
        <main className="min-h-screen p-8 bg-gray-100">
            {children}
        </main>
        </body>
        </html>
    )
}