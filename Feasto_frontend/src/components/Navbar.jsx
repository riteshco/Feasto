import {
  NavigationMenu,
  NavigationMenuList,
  NavigationMenuItem,
  NavigationMenuLink,
} from "@/components/ui/navigation-menu";
import { Button } from "@/components/ui/button";
import { ThemeToggle } from "@/components/ThemeToggle";
import { useState , useEffect } from "react";
import Cookies from "js-cookie";

export function Navbar({page , user}) {

  function LogoutUser() {
  Cookies.remove("auth_token");
  }



  const [show, setShow] = useState(true);
  const [lastScrollY, setLastScrollY] = useState(0);
   useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > lastScrollY) {
        // scrolling down
        setShow(false);
      } else {
        // scrolling up
        setShow(true);
      }
      setLastScrollY(window.scrollY);
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, [lastScrollY]);
  return (
    <header className={`bg-white dark:bg-black bg fixed top-0 left-0 z-50 w-full py-3 transition-transform duration-300 ${show ? "translate-y-0" : "-translate-y-full"}`}>
      <div className="flex items-center px-2 justify-between">
        <NavigationMenu>

          <NavigationMenuList className="flex gap-2">


            <NavigationMenuItem className="font-bold text-2xl drop-shadow-[0_0_10px_rgba(255,255,255,0.5)]">
                FEASTO
            </NavigationMenuItem>
            {page !== "Landing" ? 
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/home">
                  <Button variant="ghost">Home</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>
            :null}
            {page === "CustomerHome" || page === "AdminHome" || page === "ChefHome" ?
              <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/categories">
                  <Button variant="ghost">Categories</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>
            : null}
            {page === "CustomerHome" ? 
              <>
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/orders">
                  <Button variant="ghost">Orders</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>

            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/cart">
                  <Button variant="ghost">Cart</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>
                </>:null}
            {page === "CustomerHome" || page === "ChefHome" ? 
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/orders">
                  <Button variant="ghost">Past-Orders</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>
                :null}

            {user === "Admin" || user==="Chef" ? 
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/add-food">
                  <Button variant="ghost">Add-Food</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>
            :null}

            {user === "Admin" ? 
            <>
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/add-food">
                  <Button variant="ghost">All-Orders</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>
            
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/users">
                  <Button variant="ghost">All-Users</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>

            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/payments">
                  <Button variant="ghost">All-Payments</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>

            </>
            :null}


          </NavigationMenuList>
          
        </NavigationMenu>
        <div className="flex gap-4">
        {page !== "Landing" ? 
        <form onSubmit={LogoutUser}>
        <Button type="submit" variant="destructive_outline">Logout</Button>
        </form>
        : null}
        <ThemeToggle />
        </div>
      </div>
    </header>
  );
}
