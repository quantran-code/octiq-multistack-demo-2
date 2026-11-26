from fastapi.testclient import TestClient

from app.main import app

client = TestClient(app)


def test_health():
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "ok"}


def test_quote_without_discount():
    response = client.post("/quote", json={"unit_price_cents": 450, "quantity": 2})
    assert response.json() == {"total_cents": 900}


def test_quote_with_bulk_discount():
    response = client.post("/quote", json={"unit_price_cents": 100, "quantity": 10})
    assert response.json() == {"total_cents": 950}


def test_quote_rejects_zero_quantity():
    response = client.post("/quote", json={"unit_price_cents": 100, "quantity": 0})
    assert response.status_code == 422
