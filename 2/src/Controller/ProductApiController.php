<?php

namespace App\Controller;

use App\Repository\ProductRepository;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Attribute\Route;

final class ProductApiController extends AbstractController
{
    #[Route('/product/api', name: 'app_product_api')]
    public function index(ProductRepository $productRepository, Request $request): JsonResponse
    {
        $categoryName = $request->query->get('category');
        
        if ($categoryName) {
            // Since we don't have a strict category relationship or a join defined here yet,
            // let's assume we want to query by a field or need to implement a basic search.
            // If the Product entity doesn't have a category relationship mapped, we might need to update the query.
            // For now, let's find all and just return an array. Let's see how Product is structured.
        }

        $products = $productRepository->findAll();

        $data = [];
        foreach ($products as $product) {
            $data[] = [
                'id' => $product->getId(),
                'name' => $product->getName(),
                'price' => $product->getPrice(),
                'description' => $product->getDescription(),
                'return' => $product->getReturn(),
            ];
        }

        return $this->json($data);
    }
}
