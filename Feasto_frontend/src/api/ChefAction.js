export async function DeliverOrderAPICall(OrderId) {
    try {
    const res = await fetch(`http://localhost:3000/api/order-done/${OrderId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to deliver the order");
    return `Delivered Order #${OrderId} to the customer!`;
  } catch (err) {
    return err.message;
  }
}