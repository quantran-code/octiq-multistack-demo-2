import express from 'express'

export const SERVICES = [
  { name: 'catalog-go', port: 3001 },
  { name: 'orders-java', port: 3002 },
  { name: 'pricing-python', port: 3003 },
  { name: 'notify-ruby', port: 3004 }
]

export const SERVICE_INFO: { service: string; runtime: string; endpoints: string[] } = {
  service: 'gateway-ts',
  runtime: 'typescript/express',
  endpoints: ['/health', '/services', '/info']
}

export function createApp(): express.Express {
  const app = express()
  app.use(express.json())

  app.get('/health', (_req, res) => {
    res.json({ status: 'ok' })
  })

  app.get('/services', (_req, res) => {
    res.json({ services: SERVICES })
  })

  app.get('/info', (_req, res) => {
    res.json(SERVICE_INFO)
  })

  return app
}
