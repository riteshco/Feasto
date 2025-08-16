import { Navbar } from "@/components/Navbar"
import MainImage from "@/assets/food_home_image.jpg"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import { GetProducts } from "@/api/FetchAPI"
import { useEffect, useState } from "react"
import { AddToCartAPICall } from "@/api/Cart"
import { useNavigate } from "react-router-dom"
import { getUserFromToken } from "@/utils/Auth"
import { DeleteProductAPICall } from "@/api/Product"
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { Toaster } from "sonner"

export function CategoriesPage() {
    const user = getUserFromToken()

    const [products, setProducts] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const [quantities, setQuantities] = useState({});
    const navigate = useNavigate();

    const [selectedCategory, setSelectedCategory] = useState("all");
    
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
    
    const uniqueCategories = products
    ? [...new Set(products.map((product) => product.category))] : [];

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

    async function addToCart(productId) {
        let qty = quantities[productId] || 1;
        await AddToCartAPICall(productId, qty);
    }

    async function addOneToCart(productId) {
        await AddToCartAPICall(productId, 1);
        navigate("/cart")
    }

    async function AskToDelete(productID) {
        await DeleteProductAPICall(productID)
        const prods = await GetProducts();
        setProducts(prods);
    }

    const filteredProducts =
    selectedCategory === "all"
      ? products
      : products.filter((p) => p.category === selectedCategory);

    return (
        <>
            {user.user_role === "admin"
                ?
                <Navbar page="CategoriesPage" user="admin"/>
                :
                <Navbar page="CategoriesPage" />
            }
            <div className="relative w-full h-96 mt-16">
                <Toaster position="top-center"/>
                <div className="Main_image h-full flex justify-center">
                    <img
                        src={MainImage}
                        alt="Main home image"
                        className="w-3/4 h-full rounded-3xl object-cover"
                    />
                </div>

                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Categories!!</h1>
                    <form action="" className="flex justify-center gap-4 w-full">
                        <Select value={selectedCategory} onValueChange={(value) => setSelectedCategory(value)}>
                            <SelectTrigger className="w-1/2">
                                <SelectValue placeholder="Select a Category" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectGroup>
                                    <SelectLabel>Categories</SelectLabel>
                                    <SelectItem value="all">All</SelectItem>
                                    {uniqueCategories.map((cat) => (
                                        <SelectItem key={cat} value={cat}>{cat}</SelectItem>
                                    ))}
                                </SelectGroup>
                            </SelectContent>
                        </Select>
                    </form>
                </div>
            </div>

            <div className="flex flex-col items-center cards w-full mt-8 gap-8">
                {filteredProducts.map((product) => (

                    <Card className="w-3/4 flex">
                        <div className="cardimage w-1/4 flex justify-center items-center px-4">
                            <img src={product.image_url.String} alt="Product_image" className="w-full rounded-2xl object-fill" />
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
                                {user.user_role === "customer" ? 
                                <div className="flex flex-col gap-6">
                                    <div className="grid gap-2">
                                        <Label htmlFor={`qty-${product.id}`}>Quantity:</Label>
                                        <Input
                                        id={`qty-${product.id}`}
                                        type="number"
                                        min={1}
                                        max={1000}
                                        required
                                        value={quantities[product.id] || 1}
                                        onChange={(e) => handleQuantityChange(product.id, e.target.value)}
                                        />
                                        <Button onClick={() => addToCart(product.id)} variant="outline">Add to Cart</Button>
                                        </div>
                                    <div className="grid gap-2">
                                        <Button onClick={() => addOneToCart(product.id)}>Order Now</Button>
                                    </div>
                                </div>
                                    : 
                                    <Button onClick={() => AskToDelete(product.id)} variant="destructive_outline">Delete product</Button>
                                    }
                            </CardContent>
                        </div>
                    </Card>
                ))}

            </div>
        </>
    )
}