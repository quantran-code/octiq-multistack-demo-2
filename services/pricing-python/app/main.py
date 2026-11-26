from fastapi import FastAPI
from pydantic import BaseModel, Field

from app.pricing import quote_cents

app = FastAPI(title="pricing-python")


class QuoteRequest(BaseModel):
    unit_price_cents: int = Field(ge=0)
    quantity: int = Field(ge=1)


@app.get("/health")
def health() -> dict[str, str]:
    return {"status": "ok"}


@app.post("/quote")
def quote(request: QuoteRequest) -> dict[str, int]:
    return {"total_cents": quote_cents(request.unit_price_cents, request.quantity)}
