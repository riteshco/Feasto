import { Navbar } from "@/components/Navbar"
import MainImage from "@/assets/food_home_image.jpg"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import { GetProducts } from "@/api/fetchAPI"
import { useEffect, useState } from "react"
import { AddToCartAPICall } from "@/api/Cart"
import { useNavigate } from "react-router-dom"
import { toast, Toaster } from "sonner"
import { QuantityCounter } from "@/components/QuantityCounter"

export function CustomerHome() {

    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);



    const [searchTerm, setSearchTerm] = useState("");


    const [quantities, setQuantities] = useState({});
    const navigate = useNavigate();

    useEffect(() => {
        async function fetchProducts() {
            try {
                const prods = await GetProducts();
                setProducts(prods);
            } catch (err) {
                setError(err.message || "Failed to load products");
            } finally {
                setLoading(false);
            }
        }
        fetchProducts();

    }, []);

    const filteredProducts = products.filter((p) =>
        p.product_name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        p.category.toLowerCase().includes(searchTerm.toLowerCase())
    );

    if (loading) return <div>Loading products...</div>;
    if (error) return <div>Error: {error}</div>;

    function handleQuantityChange(productId, value) {
        let qty = Number(value);
        if (isNaN(qty) || qty < 1) qty = 1;
        if (qty > 1000) qty = 1000;

        setQuantities((prev) => ({
            ...prev,
            [productId]: qty,
        }));
    }

    const handleAddToCart = async (productID , productName) => {
    await AddToCartAPICall(productID, quantities[productID]);
    toast.success(`${productName} (x${quantities[productID]}) added to cart`);
  };

    async function addOneToCart(productId) {
        await AddToCartAPICall(productId, 1);
        navigate("/cart")
    }

    return (
        <>
            <Navbar page="CustomerHome" />
            <div className="relative w-full h-96 mt-16">
                <Toaster position="top-center" />
                <div className="Main_image h-full flex justify-center">
                    <img
                        src={MainImage}
                        alt="Main home image"
                        className="w-3/4 h-full rounded-3xl object-cover"
                    />
                </div>

                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Welcome to Feasto</h1>
                    <form action="" className="flex justify-center gap-4 w-full">
                        <Input
                            value={searchTerm}
                            onChange={(e) => setSearchTerm(e.target.value)}
                            className="w-1/2 font-extrabold text-2xl bg-white/15" placeholder="Search products..." />
                    </form>
                </div>
            </div>

            <div className="flex flex-col items-center cards w-full mt-8 gap-8">
                {console.log(products)}
                {filteredProducts ? filteredProducts.map((product) => (

                    <Card className="w-3/4 flex">
                        <div className="cardimage w-1/4 flex justify-center items-center px-4">
                            <img src={product.image_url.String} alt="Product_image" className="w-full rounded-2xl" />
                        </div>
                        <div className="cardinfo w-3/4">

                            <CardHeader>
                                <CardTitle className="text-3xl">{product.product_name}</CardTitle>
                                <CardDescription>
                                    Category : "{product.category}"
                                </CardDescription>
                                Price : ${product.price}
                            </CardHeader>
                            <CardContent>
                                <div className="flex flex-col gap-6">
                                    <div className="grid gap-2">
                                        <Label htmlFor={`qty-${product.id}`}>Quantity:</Label>
                                        <div className="flex gap-4">
                                        <QuantityCounter
                                            initialQty={1}
                                            onChange={(newQty) => handleQuantityChange(product.id , newQty)}
                                            />

                                        <Button
                                            onClick={() => {handleAddToCart(product.id , product.product_name)}}
                                            variant="outline"
                                            className="w-1/4  hover:bg-red-600 text-white"
                                            >
                                            Add to Cart
                                        </Button>
                                            </div>
                                    </div>
                                    <div className="grid gap-2">
                                        <Button onClick={() => addOneToCart(product.id)}>Order Now</Button>
                                    </div>
                                </div>
                            </CardContent>
                        </div>
                    </Card>
                ))
                    : nil}

            </div>
        </>
    )
}