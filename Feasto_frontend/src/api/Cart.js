import { toast } from "sonner";
import { API_BASE_URL } from "./Config";

export async function AddToCartAPICall(productId, quantity) {
  try {
    const res = await fetch(`${API_BASE_URL}/add-to-cart/${productId}/${quantity}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    
    if (!res.ok) {const data = await res.json() ; toast.error(data.message ||"Failed to add to cart"); return};
    toast.success(`Added ${quantity} item(s) of Product with id #${productId} to cart!`)
  } catch (err) {
    toast.error(err.message)
  }
}

export async function RemoveFromCart(OrderItemId) {
  try {
    const res = await fetch(`${API_BASE_URL}/remove-from-cart/${OrderItemId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    
    if (!res.ok) {const data = await res.json() ;toast.error(data.message ||"Failed to remove to cart"); return};
    toast.success(`Successfully removed to cart!`)
  } catch (error) {
    toast.error(err.message)
  }
}

export async function fetchCart() {
  try {

    const res = await fetch(`${API_BASE_URL}/cartItems`, {
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
  catch (error) {
    toast.error(error.message)
  }
}

export async function PlaceOrderAPICall(credentials) {
  try {
    const res = await fetch(`${API_BASE_URL}/cart/order`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(credentials),
      credentials: "include"
    });
    
    if (!res.ok) {const data = await res.json(); toast.error(data.message || "Failed to order"); return };
    toast.success(`Successfully order placed!`)
  } catch (error) {
    toast.error(error.message)
  }
}