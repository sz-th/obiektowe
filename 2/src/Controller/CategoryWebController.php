<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

final class CategoryWebController extends AbstractController
{
    #[Route('/category/web', name: 'app_category_web')]
    public function index(): Response
    {
        return $this->render('category_web/index.html.twig', [
            'controller_name' => 'CategoryWebController',
        ]);
    }
}
