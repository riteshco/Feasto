import { API_BASE_URL } from "./Config";

export async function AddFoodAPICall(credentials) {
  try {
    const res = await fetch(`${API_BASE_URL}/add-food`, {
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