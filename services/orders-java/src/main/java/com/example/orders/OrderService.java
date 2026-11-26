package com.example.orders;

import java.util.List;
import java.util.concurrent.CopyOnWriteArrayList;
import java.util.concurrent.atomic.AtomicLong;
import org.springframework.stereotype.Service;

@Service
public class OrderService {
    private final List<Order> orders = new CopyOnWriteArrayList<>();
    private final AtomicLong ids = new AtomicLong();

    public List<Order> list() {
        return List.copyOf(orders);
    }

    public Order create(String productId, int quantity) {
        if (productId == null || productId.isBlank()) {
            throw new IllegalArgumentException("productId is required");
        }
        if (quantity < 1) {
            throw new IllegalArgumentException("quantity must be at least 1");
        }
        Order order = new Order(ids.incrementAndGet(), productId, quantity);
        orders.add(order);
        return order;
    }
}
