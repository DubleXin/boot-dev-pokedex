package api

import "github.com/DubleXin/pokedexcli/pokecache"

type PokeApi struct {
	lastMapResponse *MapResponse
	Cache           *pokecache.Cache
	caughtPokemons  map[string]Pokemon
}

func (p *PokeApi) AddCaughtPokemon(pokemon Pokemon) {
	if p.caughtPokemons == nil {
		p.caughtPokemons = make(map[string]Pokemon)
	}

	p.caughtPokemons[pokemon.Name] = pokemon
}

func (p *PokeApi) GetCaughtPokemon(name string) (Pokemon, bool) {
	pokemon, ok := p.caughtPokemons[name]
	return pokemon, ok
}

func (m *PokeApi) GetCaughtPokemonNames() []string {
	names := make([]string, 0, len(m.caughtPokemons))

	for name := range m.caughtPokemons {
		names = append(names, name)
	}

	return names
}
