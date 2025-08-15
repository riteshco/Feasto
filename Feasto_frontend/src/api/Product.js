import { toast } from "sonner";
import { API_BASE_URL } from "./Config";

export async function AddFoodAPICall(credentials) {
  try {
    const res = await fetch(`${API_BASE_URL}/add-food`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(credentials),
            credentials: "include"
        });
    if (!res.ok) {const data = await res.json();toast.error(data.message || "Failed to add food to the menu"); return }
    toast.success(`Successfully Added to the menu!`)
  } catch (error) {
    toast.error(error.message)
  }
}