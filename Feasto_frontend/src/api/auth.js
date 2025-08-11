export async function loginUser(credentials) {
  const res = await fetch("http://localhost:3000/api/auth", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },

    body: JSON.stringify(credentials),
    credentials: "include", // important if you want cookies (JWT) to be set
  });

  if (!res.ok) {
    throw new Error("Login failed");
  }

  return res.json();
}

export const setCookie = (name, value) => {
  let expires = "";
    const date = new Date();
    date.setTime(date.getTime() + 60 * 60 * 1000);
    expires = "; expires=" + date.toUTCString();
  document.cookie = `${name}=${value || ""}${expires}; path=/`;
};
