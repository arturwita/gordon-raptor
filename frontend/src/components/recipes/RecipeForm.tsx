import { useEffect, useState, type FC, type FormEvent } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Plus, Trash2 } from "lucide-react";
import type { RecipeDto, RecipeInput } from "@/types/recipe";

interface RecipeFormProps {
  open: boolean;
  initialData?: RecipeDto;
  onSubmit: (data: RecipeInput) => void;
  onCancel: () => void;
  submitting?: boolean;
}

export const RecipeForm: FC<RecipeFormProps> = ({
  open,
  initialData,
  onSubmit,
  onCancel,
  submitting = false,
}) => {
  const [name, setName] = useState(initialData?.name ?? "");
  const [description, setDescription] = useState(
    initialData?.description ?? ""
  );
  const [picture, setPicture] = useState(initialData?.picture ?? "");

  const [ingredients, setIngredients] = useState<
    { name: string; quantity: string }[]
  >(
    initialData
      ? Object.entries(initialData.ingredients).map(([name, quantity]) => ({
          name,
          quantity,
        }))
      : [{ name: "", quantity: "" }]
  );

  const handleAddIngredient = () =>
    setIngredients([...ingredients, { name: "", quantity: "" }]);

  const handleRemoveIngredient = (index: number) => {
    if (ingredients.length === 1) return;
    setIngredients((prev) => prev.filter((_, i) => i !== index));
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();

    const filteredIngredients = ingredients
      .filter((i) => i.name.trim() && i.quantity.trim())
      .reduce((acc, i) => ({ ...acc, [i.name]: i.quantity }), {});

    const sanitizedDescription = description.trim();
    const sanitizedPicture = picture.trim();

    const payload: RecipeInput = {
      name: name.trim(),
      ...(sanitizedDescription && { description: sanitizedDescription }),
      ...(sanitizedPicture && { picture: sanitizedPicture }),
      ingredients: filteredIngredients,
    };

    onSubmit(payload);
  };

  useEffect(() => {
    if (initialData) {
      setName(initialData.name ?? "");
      setDescription(initialData.description ?? "");
      setPicture(initialData.picture ?? "");
      setIngredients(
        Object.entries(initialData.ingredients).map(([name, quantity]) => ({
          name,
          quantity,
        }))
      );
    } else {
      setName("");
      setDescription("");
      setPicture("");
      setIngredients([{ name: "", quantity: "" }]);
    }
  }, [initialData]);

  return (
    <Dialog open={open} onOpenChange={onCancel}>
      <DialogContent className="max-w-lg">
        <button
          onClick={onCancel}
          className="absolute top-4 right-4 text-gray-500 hover:text-gray-700"
        ></button>

        <DialogHeader>
          <DialogTitle>
            {initialData ? "Edit Recipe" : "Add Recipe"}
          </DialogTitle>
          <DialogDescription />
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4 mt-2">
          <div>
            <label className="block mb-1 font-medium">
              Name <span className="text-red-500">*</span>
            </label>
            <Input
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>

          <div>
            <label className="block mb-1 font-medium">Description</label>
            <Textarea
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />
          </div>

          <div>
            <label className="block mb-1 font-medium">Picture URL</label>
            <Input
              value={picture}
              onChange={(e) => setPicture(e.target.value)}
              placeholder="https://example.com/image.jpg"
            />
          </div>

          <div>
            <label className="block mb-2 font-medium">
              Ingredients <span className="text-red-500">*</span>
            </label>

            <div className="space-y-2">
              {ingredients.map((ingredient, index) => (
                <div key={index} className="flex gap-2">
                  <Input
                    placeholder="Name"
                    value={ingredient.name}
                    onChange={(e) => {
                      const copy = [...ingredients];
                      copy[index].name = e.target.value;
                      setIngredients(copy);
                    }}
                    required
                  />
                  <Input
                    placeholder="Quantity"
                    value={ingredient.quantity}
                    onChange={(e) => {
                      const copy = [...ingredients];
                      copy[index].quantity = e.target.value;
                      setIngredients(copy);
                    }}
                    required
                  />

                  <Button
                    type="button"
                    variant="ghost"
                    onClick={() => handleRemoveIngredient(index)}
                    disabled={ingredients.length === 1}
                  >
                    <Trash2
                      className={`h-4 w-4 ${
                        ingredients.length === 1
                          ? "opacity-40 cursor-not-allowed"
                          : "text-red-500"
                      }`}
                    />
                  </Button>
                </div>
              ))}
            </div>

            <Button
              type="button"
              variant="outline"
              size="sm"
              className="mt-2"
              onClick={handleAddIngredient}
            >
              <Plus className="h-4 w-4 mr-1" /> Add Ingredient
            </Button>
          </div>

          <div className="flex justify-end gap-2 pt-4 border-t dark:border-gray-700">
            <Button type="button" variant="outline" onClick={onCancel}>
              Cancel
            </Button>
            <Button type="submit" disabled={submitting}>
              {submitting ? "Saving..." : "Save"}
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
};
