import { API_BASE_URL } from "./Config";

export async function PaymentDoneAPICall(PaymentId) {
  try {
    const res = await fetch(`${API_BASE_URL}/payment-done/${PaymentId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to complete payment!");
    let message = `Payment complete!`
    return message

  } catch (err) {
    return err.message
  }
}