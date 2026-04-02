<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

final class ProductWebController extends AbstractController
{
    #[Route('/product/web', name: 'app_product_web')]
    public function index(): Response
    {
        return $this->render('product_web/index.html.twig', [
            'controller_name' => 'ProductWebController',
        ]);
    }
}
