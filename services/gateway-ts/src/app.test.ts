import request from 'supertest'
import { describe, expect, it } from 'vitest'
import { createApp, SERVICE_INFO, SERVICES } from './app'

describe('gateway', () => {
  it('reports health', async () => {
    const res = await request(createApp()).get('/health')
    expect(res.status).toBe(200)
    expect(res.body).toEqual({ status: 'ok' })
  })

  it('lists the other services', async () => {
    const res = await request(createApp()).get('/services')
    expect(res.status).toBe(200)
    expect(res.body.services).toEqual(SERVICES)
  })

  it('reports service info', async () => {
    const res = await request(createApp()).get('/info')
    expect(res.status).toBe(200)
    expect(res.body).toEqual(SERVICE_INFO)
  })
})
