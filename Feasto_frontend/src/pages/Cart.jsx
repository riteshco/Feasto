import { useEffect, useState } from "react";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Navbar } from "@/components/Navbar";
import { Button } from "@/components/ui/button";
import { RemoveFromCart , fetchCart , PlaceOrderAPICall } from "@/api/Cart";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@radix-ui/react-label";
import { Input } from "@/components/ui/input";

export function CartPage() {
  const [cartData, setCartData] = useState([]);

  useEffect(() => {
  async function AskToFetch() {
      let merged = await fetchCart();
      setCartData(merged);
  }
  AskToFetch();
  }, []);

    async function AskToRemoveFromCart(Id){
        setCartData(prev => prev.filter(item => item.id !== Id));
        await RemoveFromCart(Id);
        let merged = await fetchCart();
        setCartData(merged);
    }

    const PlaceCartOrder = async (e)=>{
        e.preventDefault(); //Stops form from reloading
        const table_number = parseInt(e.target.table_number.value , 10)
        const instructions = e.target.instructions.value

        await PlaceOrderAPICall({table_number , instructions})
        let merged = await fetchCart();
        setCartData(merged);
    }

    return (
        <>
            <Navbar page="CartPage" />
            <div className="relative w-full h-96 mt-16">
                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Cart!!</h1>
                </div>
            </div>

            <div className="flex flex-col items-center cards w-full mt-8 gap-8">
                {cartData ? cartData.map((cartItem) => (

                    <Card className="w-3/4 flex">
                        <div className="cardinfo w-3/4">
                            <CardHeader>
                                <CardTitle className="text-3xl">{cartItem.product_name}</CardTitle>
                                <div className="quantity">
                                Quantity : {cartItem.quantity}
                                </div>
                                <div className="price">
                                Price : {cartItem.price}
                                </div>
                            </CardHeader>
                            <CardContent>
                                <div className="w-full">
                                        <Button onClick={() => AskToRemoveFromCart(cartItem.orderId)} variant="destructive_outline">Remove From Cart</Button>
                                </div>
                            </CardContent>
                        </div>
                    </Card>
                ))
            : null}
            <form className="flex flex-col w-3/4 gap-4" onSubmit={PlaceCartOrder}>
                <Label htmlFor="table_number">Table Number:</Label>
                <Input id="table_number"
                type="number"
                min={1}
                max={1000}
                required
                />
                <Label htmlFor="instructions">Instructions:</Label>
                <Textarea id="instructions" placeholder="Example : Please make it more salty!!........" required />
                <Button className="w-full" >Place Your Order</Button>
            </form>
            </div>
        </>
    )
}