import { useEffect, useState } from "react"
import { Button } from "./ui/button"

export function ThemeToggle() {
  const [isDark, setIsDark] = useState(() => {
    // On first load, check localStorage or system preference
    if (typeof window !== "undefined") {
      const storedTheme = localStorage.getItem("theme")
      if (storedTheme) return storedTheme === "dark"
      return window.matchMedia("(prefers-color-scheme: dark)").matches
    }
    return false
  })

  useEffect(() => {
    if (isDark) {
      document.documentElement.classList.add("dark")
      localStorage.setItem("theme", "dark")
    } else {
      document.documentElement.classList.remove("dark")
      localStorage.setItem("theme", "light")
    }
  }, [isDark])

  return (
    <Button
      className="px-4 py-2 rounded-md border"
      onClick={() => setIsDark(!isDark)}
    >
      {isDark ? "☀ Light Mode" : "🌙 Dark Mode"}
    </Button>
  )
}
