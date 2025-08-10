import {
  NavigationMenu,
  NavigationMenuList,
  NavigationMenuItem,
  NavigationMenuLink,
} from "@/components/ui/navigation-menu";
import { Button } from "@/components/ui/button";
import { ThemeToggle } from "@/components/ThemeToggle";

export function Navbar({page , user}) {
  return (
    <header className="w-full py-3">
      <div className="flex items-center px-2 justify-between">
        <NavigationMenu>

          <NavigationMenuList className="flex gap-2">


            <NavigationMenuItem className="font-bold text-2xl drop-shadow-[0_0_10px_rgba(255,255,255,0.5)]">
                FEASTO
            </NavigationMenuItem>
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <a href="/home">
                  <Button variant="ghost">Home</Button>
                </a>
              </NavigationMenuLink>
            </NavigationMenuItem>
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
        <Button variant="destructive_outline">Logout</Button>
        <ThemeToggle />
        </div>
      </div>
    </header>
  );
}
