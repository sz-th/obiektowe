<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

final class OrderWebController extends AbstractController
{
    #[Route('/order/web', name: 'app_order_web')]
    public function index(): Response
    {
        return $this->render('order_web/index.html.twig', [
            'controller_name' => 'OrderWebController',
        ]);
    }
}
