import { Navbar } from "@/components/Navbar";
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { AddFoodAPICall } from "@/api/Product";
import { getUserFromToken } from "@/utils/Auth";
import { Check, ChevronsUpDown } from "lucide-react"
import { Toaster , toast } from "sonner";

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
import { useEffect, useState } from "react";
import { GetProducts } from "@/api/FetchAPI";

export function AddFoodPage() {

    const [open, setOpen] = useState(false)
    const [value, setValue] = useState("")
    const [categories, setCategory] = useState([])

    useEffect(() => {
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
    }, [])

    const AskToAddFood = async (e) => {
        e.preventDefault();
        const product_name = e.target.product_name.value
        const price = parseFloat(e.target.price.value)
        const category = e.target.category.value
        const image_url = e.target.image_url.value

        const message = await AddFoodAPICall({ product_name, price, category, image_url })
        toast(message, {
                action: {
                    label: "Ok",
                },
        })
        e.target.product_name.value = ""
        e.target.price.value = ""
        e.target.category.value = ""
        e.target.image_url.value = ""
    }

    const user = getUserFromToken()

    return (
        <>
            {user.user_role === "admin"
                ?
                <Navbar page="AddFoodPage" user="admin" />
                :
                <Navbar page="AddFoodPage" />
            }
            <div className="body flex flex-col items-center justify-center mt-28 gap-12">
                <Toaster position="top-center"/>
                <div className="title">
                    <div className="text-5xl font-extrabold">Add-Food</div>
                </div>
                <div className="card w-1/2">

                    <Card className="w-full">
                        <CardHeader>
                            <CardTitle>Add new Food item to the Menu!</CardTitle>
                            <CardDescription>
                                Enter all the details given below to add it to our Restaurant's great Menu!
                            </CardDescription>
                        </CardHeader>
                        <CardContent>
                            <form onSubmit={AskToAddFood}>
                                <div className="flex flex-col gap-6">
                                    <div className="grid gap-2">
                                        <Label htmlFor="product_name">Food Name:</Label>
                                        <Input
                                            id="product_name"
                                            type="text"
                                            placeholder="Name for new menu item......"
                                            required
                                        />
                                    </div>
                                    <div className="grid gap-2">
                                        <Label htmlFor="price">Price:</Label>
                                        <Input
                                            id="price"
                                            type="number"
                                            min={0.01}
                                            max={1000.00}
                                            step={0.01}
                                            placeholder="$...."
                                            required
                                        />
                                    </div>
                                    <div className="grid gap-2">
                                        <Label htmlFor="category">Category:</Label>
                                        <Popover open={open} onOpenChange={setOpen}>
                                            <PopoverTrigger asChild>
                                                <Button
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
                                            <PopoverContent className="w-[250px] p-0">
                                                <Command>
                                                    <CommandInput placeholder="Search or add new category..." className="h-9" />
                                                    <CommandList>
                                                        <CommandEmpty>
                                                            <Button
                                                                className="w-full h-8 text-sm"
                                                                onClick={() => {
                                                                    // Example: set typed value as new category
                                                                    const newCategory = document.querySelector(
                                                                        '[placeholder="Search or add new category..."]'
                                                                    ).value.trim();

                                                                    if (newCategory && !categories.some((c) => c.value === newCategory)) {
                                                                        const newCatObj = { value: newCategory, label: newCategory };
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
                                                                    onSelect={(currentValue) => {
                                                                        setValue(currentValue === value ? "" : currentValue);
                                                                        setOpen(false);
                                                                    }}
                                                                >
                                                                    {Category.label}
                                                                    <Check
                                                                        className={cn(
                                                                            "ml-auto",
                                                                            value === Category.value ? "opacity-100" : "opacity-0"
                                                                        )}
                                                                    />
                                                                </CommandItem>
                                                            ))}
                                                        </CommandGroup>
                                                    </CommandList>
                                                </Command>
                                            </PopoverContent>
                                        </Popover>
                                        <input type="hidden" name="category" value={value} required />
                                    </div>

                                    <div className="grid gap-2">
                                        <Label htmlFor="image_url">URL for the product image:</Label>
                                        <Input
                                            id="image_url"
                                            type="text"
                                            placeholder="Leave empty for null value"
                                        />
                                    </div>
                                </div>
                                <div className="flex justify-center mt-8">
                                    <Button type="submit" className="w-1/2">
                                        Add the Food item
                                    </Button>
                                </div>
                            </form>
                        </CardContent>
                    </Card>
                </div>
            </div>
        </>
    )
}