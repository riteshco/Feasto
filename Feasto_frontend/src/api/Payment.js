import { toast } from "sonner";
import { API_BASE_URL } from "./Config";

export async function PaymentDoneAPICall(PaymentId) {
  try {
    const res = await fetch(`${API_BASE_URL}/payment-done/${PaymentId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) {const data = await res.json();toast.error(data.message || "Failed to complete payment!"); return};
    toast.success(`Payment complete!`)

  } catch (err) {
    toast.error(err.message)
  }
}