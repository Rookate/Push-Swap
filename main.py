import random
import string

def generate_sequence():
    # Choisir aléatoirement de générer une séquence correcte ou incorrecte
    is_correct = random.choice([True, False])
    
    if is_correct:
        # Générer une suite correcte de 5 chiffres compris entre 0 et 100
        sequence = [random.randint(0, 100) for _ in range(5)]
        print(' '.join(map(str, sequence)))
    else:
        # Générer une suite incorrecte avec au moins une lettre et le reste des nombres
        sequence = [random.randint(0, 100) for _ in range(4)]
        sequence.append(random.choice(string.ascii_letters))  # Ajouter au moins une lettre
        random.shuffle(sequence)  # Mélanger la séquence pour éviter que la lettre soit toujours en dernière position
        print(' '.join(map(str, sequence)))

# Appeler la fonction pour générer une séquence
generate_sequence()