import { Navbar } from "@/components/Navbar"
import MainImage from "@/assets/food_home_image.jpg"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { UpdateFoodAPICall } from "@/api/Product"
import { cn } from "@/lib/Utils"
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
} from "@/components/ui/command"
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/popover"
import { Check, ChevronsUpDown } from "lucide-react"
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

    const [open, setOpen] = useState(false)
    const [products, setProducts] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [value, setValue] = useState("")

    const [categories, setCategory] = useState([])

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
        async function AskForFoodCategories() {
            const data = await GetProducts()
            let cats = []
            const uniquedata = [...new Map(data.map(item => [item.category, item])).values()];
            uniquedata.map((product) => {
                cats.push({ value: product.category, label: product.category })
            })
            setCategory(cats)
        }
        AskForFoodCategories();
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
    const AskToUpdateFood = async (e, id) => {
        e.preventDefault();

        const formData = new FormData(e.target);
        const data = {
            product_name: formData.get("new_product_name"),
            price: parseFloat(formData.get("new_price")),
            category: formData.get("new_category"),
            image_url: formData.get("new_image_url") || null,
        };

        await UpdateFoodAPICall(data, id)
    }

    const filteredProducts =
        selectedCategory === "all"
            ? products
            : products.filter((p) => p.category === selectedCategory);

    return (
        <>
            {user.user_role === "admin"
                ?
                <Navbar page="CategoriesPage" user="admin" />
                :
                <Navbar page="CategoriesPage" />
            }
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
                                    <div className="flex flex-col gap-4">
                                        <Button onClick={() => AskToDelete(product.id)} className="w-1/4" variant="destructive_outline">Delete product</Button>
                                        <Dialog>
                                            <DialogTrigger asChild>
                                                {/* 👇 prevent accidental submit */}
                                                <Button type="button" variant="outline" className="w-1/4">
                                                    Edit details
                                                </Button>
                                            </DialogTrigger>

                                            <DialogContent>
                                                <DialogHeader>
                                                    <DialogTitle>Edit details</DialogTitle>
                                                    <DialogDescription>
                                                        Make changes to the product here. Click save when you&apos;re done.
                                                    </DialogDescription>
                                                </DialogHeader>

                                                <form onSubmit={(e) => AskToUpdateFood(e, product.id)}>
                                                    <div className="grid gap-4">
                                                        <div className="grid gap-3">
                                                            <Label htmlFor="new_product_name">Product Name</Label>
                                                            <Input
                                                                id="new_product_name"
                                                                name="new_product_name"
                                                                defaultValue={product.product_name}
                                                                required
                                                            />
                                                        </div>

                                                        {/* Price */}
                                                        <div className="grid gap-3">
                                                            <Label htmlFor="new_price">Price</Label>
                                                            <Input
                                                                id="new_price"
                                                                name="new_price"
                                                                type="number"
                                                                min={1}
                                                                max={1000}
                                                                step={0.01}
                                                                defaultValue={product.price}
                                                                required
                                                            />
                                                        </div>

                                                        {/* Category */}
                                                        <div className="grid gap-2">
                                                            <Label htmlFor="category">Category:</Label>
                                                            <Popover open={open} onOpenChange={setOpen}>
                                                                <PopoverTrigger asChild>
                                                                    {/* 👇 must NOT submit */}
                                                                    <Button
                                                                        type="button"
                                                                        variant="outline"
                                                                        role="combobox"
                                                                        aria-expanded={open}
                                                                        className="w-[200px] justify-between"
                                                                    >
                                                                        {value
                                                                            ? categories.find((c) => c.value === value)?.label
                                                                            : "Select or Add Category..."}
                                                                        <ChevronsUpDown className="opacity-50" />
                                                                    </Button>
                                                                </PopoverTrigger>

                                                                <PopoverContent className="w-[250px] p-0 z-[9999]">
                                                                    <Command>
                                                                        <CommandInput
                                                                            placeholder="Search or add new category..."
                                                                            className="h-9"
                                                                        />
                                                                        <CommandList>
                                                                            <CommandEmpty>
                                                                                <Button
                                                                                    type="button"
                                                                                    className="w-full h-8 text-sm"
                                                                                    onClick={() => {
                                                                                        const newCategory = document.querySelector(
                                                                                            '[placeholder="Search or add new category..."]'
                                                                                        ).value.trim();

                                                                                        if (
                                                                                            newCategory &&
                                                                                            !categories.some((c) => c.value === newCategory)
                                                                                        ) {
                                                                                            const newCatObj = {
                                                                                                value: newCategory,
                                                                                                label: newCategory,
                                                                                            };
                                                                                            setCategory((prev) => [...prev, newCatObj]);
                                                                                            setValue(newCategory);
                                                                                            setOpen(false);
                                                                                        }
                                                                                    }}
                                                                                >
                                                                                    ➕ Add new category
                                                                                </Button>
                                                                            </CommandEmpty>

                                                                            <CommandGroup>
                                                                                {categories.map((Category) => (
                                                                                    <CommandItem
                                                                                        key={Category.value}
                                                                                        value={Category.value}
                                                                                        onMouseDown={(e) => e.preventDefault()}
                                                                                        onSelect={(currentValue) => {
                                                                                            setValue(currentValue === value ? "" : currentValue);
                                                                                            setOpen(false);
                                                                                        }}
                                                                                    >
                                                                                        {Category.label}
                                                                                        <Check
                                                                                            className={cn(
                                                                                                "ml-auto",
                                                                                                value === Category.value
                                                                                                    ? "opacity-100"
                                                                                                    : "opacity-0"
                                                                                            )}
                                                                                        />
                                                                                    </CommandItem>
                                                                                ))}
                                                                            </CommandGroup>
                                                                        </CommandList>
                                                                    </Command>
                                                                </PopoverContent>
                                                            </Popover>

                                                            <input type="hidden" name="new_category" value={value} required />
                                                        </div>

                                                        {/* Image URL */}
                                                        <div className="grid gap-2">
                                                            <Label htmlFor="new_image_url">URL for the product image:</Label>
                                                            <Input
                                                                id="new_image_url"
                                                                name="new_image_url"
                                                                type="text"
                                                                placeholder="Leave empty for null value"
                                                            />
                                                        </div>
                                                    </div>

                                                    <DialogFooter className="mt-8">
                                                        <DialogClose asChild>
                                                            <Button type="button" variant="outline">
                                                                Cancel
                                                            </Button>
                                                        </DialogClose>
                                                        <Button type="submit">Save changes</Button>
                                                    </DialogFooter>
                                                </form>
                                            </DialogContent>
                                        </Dialog>

                                    </div>

                                }
                            </CardContent>
                        </div>
                    </Card>
                ))}

            </div>
        </>
    )
}