export async function AddChangeRequest(role) {
  try {
    const res = await fetch(`http://localhost:3000/api/change_role_request/${role}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to request role change!");
    let message = `Request to change role sent`
    return message

  } catch (err) {
    return err.message
  }
}