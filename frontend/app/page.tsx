'use client' // 必须标记为客户端组件

import { useState } from 'react'
import { createPromiseClient } from '@bufbuild/connect'
import { createConnectTransport } from '@bufbuild/connect-web'
import { CalculatorService } from './proto/calculator/v1/calculator_connectweb.js'
import { CalculationRequest } from './proto/calculator/v1/calculator_pb.js'

export default function CalculatorPage() {
    const [num1, setNum1] = useState('')
    const [num2, setNum2] = useState('')
    const [operation, setOperation] = useState('+')
    const [result, setResult] = useState<number | null>(null)
    const [error, setError] = useState('')
    const [isLoading, setIsLoading] = useState(false)

    const calculate = async () => {
        if (!num1 || !num2) {
            setError('Please enter both numbers')
            return
        }

        setIsLoading(true)
        setError('')
        setResult(null)

        try {
            // 创建 Connect 传输层
            const transport = createConnectTransport({
                baseUrl: '/api', // 通过 Next.js rewrite 代理
            })

            // 创建客户端
            const client = createPromiseClient(CalculatorService, transport)

            // 准备请求
            const req = new CalculationRequest({
                num1: parseFloat(num1),
                num2: parseFloat(num2),
                operation,
            })

            // 调用 RPC 方法
            const response = await client.calculate(req)
            setResult(response.result)
        } catch (err) {
            setError(err instanceof Error ? err.message : 'Unknown error')
        } finally {
            setIsLoading(false)
        }
    }

    return (
        <div className="max-w-md mx-auto p-6 bg-white rounded-lg shadow-md">
            <h1 className="text-2xl font-bold mb-6 text-center">ConnectRPC Calculator</h1>

            <div className="space-y-4">
                {/* 第一个数字输入 */}
                <div>
                    <label className="block text-sm font-medium mb-1">First Number</label>
                    <input
                        type="number"
                        value={num1}
                        onChange={(e) => setNum1(e.target.value)}
                        className="w-full p-2 border rounded"
                        placeholder="Enter first number"
                    />
                </div>

                {/* 操作符选择 */}
                <div>
                    <label className="block text-sm font-medium mb-1">Operation</label>
                    <select
                        value={operation}
                        onChange={(e) => setOperation(e.target.value)}
                        className="w-full p-2 border rounded"
                    >
                        <option value="+">Addition (+)</option>
                        <option value="-">Subtraction (-)</option>
                        <option value="*">Multiplication (*)</option>
                        <option value="/">Division (/)</option>
                    </select>
                </div>

                {/* 第二个数字输入 */}
                <div>
                    <label className="block text-sm font-medium mb-1">Second Number</label>
                    <input
                        type="number"
                        value={num2}
                        onChange={(e) => setNum2(e.target.value)}
                        className="w-full p-2 border rounded"
                        placeholder="Enter second number"
                    />
                </div>

                {/* 计算按钮 */}
                <button
                    onClick={calculate}
                    disabled={isLoading}
                    className={`w-full py-2 px-4 rounded-md text-white ${
                        isLoading ? 'bg-gray-400' : 'bg-blue-600 hover:bg-blue-700'
                    }`}
                >
                    {isLoading ? 'Calculating...' : 'Calculate'}
                </button>

                {/* 结果显示 */}
                {result !== null && (
                    <div className="mt-4 p-4 bg-gray-50 rounded-md">
                        <h2 className="text-lg font-medium">Result</h2>
                        <p className="text-xl font-mono">{result}</p>
                    </div>
                )}

                {/* 错误显示 */}
                {error && (
                    <div className="mt-4 p-4 bg-red-50 text-red-600 rounded-md">
                        <p>{error}</p>
                    </div>
                )}
            </div>
        </div>
    )
}