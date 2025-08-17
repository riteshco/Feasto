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
import { Link, useNavigate } from "react-router-dom";

export function Navbar({page , user}) {

  const navigate = useNavigate()

  function LogoutUser() {
  Cookies.remove("auth_token");
  navigate("/")
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
      <div className="flex items-center px-8 justify-between">
        <NavigationMenu>

          <NavigationMenuList className="flex gap-2">
            <NavigationMenuItem className="font-bold text-2xl drop-shadow-[0_0_10px_rgba(255,255,255,1)]">
                FEASTO
            </NavigationMenuItem>
            {page !== "Landing" ? 
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                {user === "admin" ?
                // <Link to="/admin">
                //   <Button variant="ghost">Admin Panel</Button>
                // </Link>
                null
                :
                <Link to="/home">
                  <Button variant="ghost">Home</Button>
                </Link>
                }
              </NavigationMenuLink>
            </NavigationMenuItem>
            :null}
            {page === "CustomerHome" || page === "AdminDashboard" || page === "ChefHome" || user === "admin" ?
              <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/categories">
                  <Button variant="ghost">Categories</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
            : null}
            {page === "CustomerHome" || page==="ChefHome" ? 
              <>
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/orders">
                  <Button variant="ghost">Orders</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>

              <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/change-role">
                  <Button variant="ghost">Change Role</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>

            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/cart">
                  <Button variant="ghost">Cart</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
                </>:null}
            {page === "ChefHome" ? 
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/delivered-orders">
                  <Button variant="ghost">Delivered-Orders</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
                :null}
            {page === "CustomerHome" || page === "ChefHome" ? 
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/past-orders">
                  <Button variant="ghost">Past-Orders</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
                :null}

            {user === "admin" || user==="chef" ? 
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/add-food">
                  <Button variant="ghost">Add-Food</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
            :null}

            {user === "admin" ? 
            <>
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/all-orders">
                  <Button variant="ghost">All-Orders</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
            
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/users">
                  <Button variant="ghost">All-Users</Button>
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>

            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Link to="/all-payments">
                  <Button variant="ghost">All-Payments</Button>
                </Link>
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
