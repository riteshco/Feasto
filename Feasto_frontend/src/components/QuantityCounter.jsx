import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export function QuantityCounter({ onChange, initialQty = 1 }) {
  const [qty, setQty] = useState(initialQty);

  const updateQty = (newQty) => {
    if (newQty < 1) return;
    setQty(newQty);
    onChange?.(newQty);
  };

  return (
    <div className="flex items-center gap-2">
      <Button variant="outline" size="sm" onClick={() => updateQty(qty - 1)}>
        –
      </Button>
      <Input
        type="number"
        value={qty}
        onChange={(e) => updateQty(Number(e.target.value))}
        className="w-16 text-center"
        min={1}
      />
      <Button variant="outline" size="sm" onClick={() => updateQty(qty + 1)}>
        +
      </Button>
    </div>
  );
}