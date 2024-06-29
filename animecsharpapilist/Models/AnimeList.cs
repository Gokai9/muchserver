namespace animecsharpapilist.Models;

public class AnimeList
{
    public DateOnly Date { get; set; }
    public string? Title { get; set; }
    public int Episode { get; set; }
    public bool IsComplete { get; set; }
}
