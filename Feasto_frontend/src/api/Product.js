export async function AddFoodAPICall(credentials) {
  try {
    const res = await fetch("http://localhost:3000/api/add-food", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(credentials),
            credentials: "include"
        });
    if (!res.ok) throw new Error("Failed to add food to the menu");
    return `Successfully Added to the menu!`
  } catch (error) {
    return error.message
  }
}