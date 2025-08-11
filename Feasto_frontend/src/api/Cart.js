export async function AddToCartAPICall(productId , quantity) {
    try {
      const res = await fetch(`http://localhost:3000/api/add-to-cart/${productId}/${quantity}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
      if (!res.ok) throw new Error("Failed to add to cart");
      alert(`Added ${quantity} item(s) to cart!`);
    } catch (err) {
      alert(err.message);
    }
  }