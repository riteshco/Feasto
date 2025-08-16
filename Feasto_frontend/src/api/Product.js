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

export async function DeleteProductAPICall(productID) {
  try{
    const res = await fetch(`${API_BASE_URL}/delete-product/${[productID]}`, {
            method: "DELETE",
            headers: { "Content-Type": "application/json" },
            credentials: "include"
        });
    if (!res.ok) {const data = await res.json();toast.error(data.message || "Failed to add food to the menu"); return }
    toast.success(`Successfully deleted the product!`)
  } catch(error){
    toast.error(error.message)
  }
}