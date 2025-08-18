import { toast } from "sonner";
import { API_BASE_URL } from "./Config";

export async function DeliverOrderAPICall(OrderId) {
    try {
    const res = await fetch(`${API_BASE_URL}/order-done/${OrderId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    
    if (!res.ok) {const data = await res.json();toast.error(data.message || "Failed to deliver the order"); return};
    toast.success(`Delivered Order #${OrderId} to the customer!`)
  } catch (err) {
    toast.error(err.message)
  }
}

export async function StartOrderAPICall(OrderId) {
    try {
    const res = await fetch(`${API_BASE_URL}/take-order/${OrderId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    
    if (!res.ok) {const data = await res.json();toast.error(data.message || "Failed to take the order"); return};
    toast.success(`You Started Order #${OrderId} for the customer!`)
  } catch (err) {
    toast.error(err.message)
  }
}