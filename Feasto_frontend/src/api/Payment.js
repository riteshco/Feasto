export async function PaymentDoneAPICall(PaymentId) {
  try {
    const res = await fetch(`http://localhost:3000/api/payment-done/${PaymentId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to complete payment!");
    alert(`Payment complete!`);
  } catch (err) {
    alert(err.message);
  }
}