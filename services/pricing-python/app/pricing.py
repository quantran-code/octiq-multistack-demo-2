BULK_THRESHOLD = 10
BULK_DISCOUNT_PERCENT = 5


def quote_cents(unit_price_cents: int, quantity: int) -> int:
    """Total price in cents, with a discount for bulk quantities."""
    total = unit_price_cents * quantity
    if quantity >= BULK_THRESHOLD:
        total -= total * BULK_DISCOUNT_PERCENT // 100
    return total
