export async function AddToCartAPICall(productId, quantity) {
  try {
    const res = await fetch(`http://localhost:3000/api/add-to-cart/${productId}/${quantity}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to add to cart");
    return `Added ${quantity} item(s) of Product with id #${productId} to cart!`
  } catch (err) {
    return err.message
  }
}

export async function RemoveFromCart(OrderItemId) {
  try {
    const res = await fetch(`http://localhost:3000/api/remove-from-cart/${OrderItemId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to remove to cart");
    return `Successfully removed to cart!`
  } catch (error) {
    return err.message
  }
}

export async function fetchCart() {
  const res = await fetch("http://localhost:3000/api/cartItems", {
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  const data = await res.json();
  if (data.orders) {

    return data.orders.map(order => {
      const product = data.products.find(p => p.id === order.product_id);
      return {
        orderId: order.id,
        product_name: product ? product.product_name : "Unknown Product",
        quantity: order.quantity,
        price: product.price
      };
    });
  } else {
    return null
  }
}

export async function PlaceOrderAPICall(credentials) {
  try {
    const res = await fetch("http://localhost:3000/api/cart/order", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(credentials),
            credentials: "include"
        });
    if (!res.ok) throw new Error("Failed order from cart");
    return `Successfully order placed!`
  } catch (error) {
    return error.message
  }
}